PGDMP                         {            go-resep    15.2    15.2 (                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            !           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            "           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            #           1262    18678    go-resep    DATABASE     �   CREATE DATABASE "go-resep" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE "go-resep";
                postgres    false            �            1259    18824    bahan    TABLE     �   CREATE TABLE public.bahan (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    nama_bahan character varying(255),
    satuan character varying(255)
);
    DROP TABLE public.bahan;
       public         heap    postgres    false            �            1259    18823    bahan_id_seq    SEQUENCE     u   CREATE SEQUENCE public.bahan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.bahan_id_seq;
       public          postgres    false    215            $           0    0    bahan_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.bahan_id_seq OWNED BY public.bahan.id;
          public          postgres    false    214            �            1259    18855    detail_resep    TABLE     )  CREATE TABLE public.detail_resep (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    resep_id bigint NOT NULL,
    bahan_id bigint NOT NULL,
    "'jumlah'" integer,
    jumlah integer NOT NULL
);
     DROP TABLE public.detail_resep;
       public         heap    postgres    false            �            1259    18854    detail_resep_id_seq    SEQUENCE     |   CREATE SEQUENCE public.detail_resep_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.detail_resep_id_seq;
       public          postgres    false    221            %           0    0    detail_resep_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.detail_resep_id_seq OWNED BY public.detail_resep.id;
          public          postgres    false    220            �            1259    18834    kategori    TABLE     �   CREATE TABLE public.kategori (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    kategori character varying(255)
);
    DROP TABLE public.kategori;
       public         heap    postgres    false            �            1259    18833    kategori_id_seq    SEQUENCE     x   CREATE SEQUENCE public.kategori_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.kategori_id_seq;
       public          postgres    false    217            &           0    0    kategori_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.kategori_id_seq OWNED BY public.kategori.id;
          public          postgres    false    216            �            1259    18842    resep    TABLE     �   CREATE TABLE public.resep (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    nama_resep character varying(255),
    kategori_id bigint NOT NULL
);
    DROP TABLE public.resep;
       public         heap    postgres    false            �            1259    18841    resep_id_seq    SEQUENCE     u   CREATE SEQUENCE public.resep_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.resep_id_seq;
       public          postgres    false    219            '           0    0    resep_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.resep_id_seq OWNED BY public.resep.id;
          public          postgres    false    218            t           2604    18827    bahan id    DEFAULT     d   ALTER TABLE ONLY public.bahan ALTER COLUMN id SET DEFAULT nextval('public.bahan_id_seq'::regclass);
 7   ALTER TABLE public.bahan ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            w           2604    18858    detail_resep id    DEFAULT     r   ALTER TABLE ONLY public.detail_resep ALTER COLUMN id SET DEFAULT nextval('public.detail_resep_id_seq'::regclass);
 >   ALTER TABLE public.detail_resep ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    220    221            u           2604    18837    kategori id    DEFAULT     j   ALTER TABLE ONLY public.kategori ALTER COLUMN id SET DEFAULT nextval('public.kategori_id_seq'::regclass);
 :   ALTER TABLE public.kategori ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            v           2604    18845    resep id    DEFAULT     d   ALTER TABLE ONLY public.resep ALTER COLUMN id SET DEFAULT nextval('public.resep_id_seq'::regclass);
 7   ALTER TABLE public.resep ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    219    218    219                      0    18824    bahan 
   TABLE DATA           [   COPY public.bahan (id, created_at, updated_at, deleted_at, nama_bahan, satuan) FROM stdin;
    public          postgres    false    215   �,                 0    18855    detail_resep 
   TABLE DATA           v   COPY public.detail_resep (id, created_at, updated_at, deleted_at, resep_id, bahan_id, "'jumlah'", jumlah) FROM stdin;
    public          postgres    false    221   �-                 0    18834    kategori 
   TABLE DATA           T   COPY public.kategori (id, created_at, updated_at, deleted_at, kategori) FROM stdin;
    public          postgres    false    217   .                 0    18842    resep 
   TABLE DATA           `   COPY public.resep (id, created_at, updated_at, deleted_at, nama_resep, kategori_id) FROM stdin;
    public          postgres    false    219   �.       (           0    0    bahan_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.bahan_id_seq', 9, true);
          public          postgres    false    214            )           0    0    detail_resep_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.detail_resep_id_seq', 3, true);
          public          postgres    false    220            *           0    0    kategori_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.kategori_id_seq', 4, true);
          public          postgres    false    216            +           0    0    resep_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.resep_id_seq', 1, true);
          public          postgres    false    218            y           2606    18831    bahan bahan_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.bahan
    ADD CONSTRAINT bahan_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.bahan DROP CONSTRAINT bahan_pkey;
       public            postgres    false    215            �           2606    18860    detail_resep detail_resep_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.detail_resep
    ADD CONSTRAINT detail_resep_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.detail_resep DROP CONSTRAINT detail_resep_pkey;
       public            postgres    false    221            }           2606    18839    kategori kategori_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.kategori
    ADD CONSTRAINT kategori_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.kategori DROP CONSTRAINT kategori_pkey;
       public            postgres    false    217            �           2606    18847    resep resep_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.resep
    ADD CONSTRAINT resep_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.resep DROP CONSTRAINT resep_pkey;
       public            postgres    false    219            z           1259    18832    idx_bahan_deleted_at    INDEX     L   CREATE INDEX idx_bahan_deleted_at ON public.bahan USING btree (deleted_at);
 (   DROP INDEX public.idx_bahan_deleted_at;
       public            postgres    false    215            �           1259    18871    idx_detail_resep_deleted_at    INDEX     Z   CREATE INDEX idx_detail_resep_deleted_at ON public.detail_resep USING btree (deleted_at);
 /   DROP INDEX public.idx_detail_resep_deleted_at;
       public            postgres    false    221            {           1259    18840    idx_kategori_deleted_at    INDEX     R   CREATE INDEX idx_kategori_deleted_at ON public.kategori USING btree (deleted_at);
 +   DROP INDEX public.idx_kategori_deleted_at;
       public            postgres    false    217            ~           1259    18853    idx_resep_deleted_at    INDEX     L   CREATE INDEX idx_resep_deleted_at ON public.resep USING btree (deleted_at);
 (   DROP INDEX public.idx_resep_deleted_at;
       public            postgres    false    219            �           2606    18861 "   detail_resep fk_detail_resep_bahan    FK CONSTRAINT     �   ALTER TABLE ONLY public.detail_resep
    ADD CONSTRAINT fk_detail_resep_bahan FOREIGN KEY (bahan_id) REFERENCES public.bahan(id);
 L   ALTER TABLE ONLY public.detail_resep DROP CONSTRAINT fk_detail_resep_bahan;
       public          postgres    false    215    221    3193            �           2606    18866 "   detail_resep fk_detail_resep_resep    FK CONSTRAINT     �   ALTER TABLE ONLY public.detail_resep
    ADD CONSTRAINT fk_detail_resep_resep FOREIGN KEY (resep_id) REFERENCES public.resep(id);
 L   ALTER TABLE ONLY public.detail_resep DROP CONSTRAINT fk_detail_resep_resep;
       public          postgres    false    221    3200    219            �           2606    18873 "   detail_resep fk_resep_detail_resep    FK CONSTRAINT     �   ALTER TABLE ONLY public.detail_resep
    ADD CONSTRAINT fk_resep_detail_resep FOREIGN KEY (resep_id) REFERENCES public.resep(id);
 L   ALTER TABLE ONLY public.detail_resep DROP CONSTRAINT fk_resep_detail_resep;
       public          postgres    false    221    3200    219            �           2606    18848    resep fk_resep_kategori    FK CONSTRAINT     }   ALTER TABLE ONLY public.resep
    ADD CONSTRAINT fk_resep_kategori FOREIGN KEY (kategori_id) REFERENCES public.kategori(id);
 A   ALTER TABLE ONLY public.resep DROP CONSTRAINT fk_resep_kategori;
       public          postgres    false    217    219    3197               �   x���;�0Dk����ݵ�ϖ�.�&-�(� �!B�=.@(�j�7�4�QV���]������h�#�^tÔ{� �BR���g�䚸U���YF���o�.ͣ���4^d�ĩb
�F��b�*-�m��Wm���pe��Q��-+ܪtA����<�_���	կF�         Q   x�3�4202�5��50R00�2��26׳4�454�50' ��i�i	������Y�)P�����1�IXd�&�)�=... � F         �   x����
1D��W�vg������hy��ŉ��\A�f���c4���w�%�����b�j��N�0���F�,I�T�T���p�[���K$p�l&����y������E�IE���s����Xh�D���-b         ?   x�3�4202�5��50R00�2��26׳40�05�50' ���X��������i����� ��S     