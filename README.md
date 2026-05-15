**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-05-15 18:30:42 UTC（2026-05-16 02:30:42 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 639ms | ✓ 778ms | ✓ 621ms | ✓ 847ms | ✓ 695ms | http |
| 137.59.47.73:3128 | ✓ 1617ms | ✓ 1178ms | ✓ 1050ms | ✓ 1373ms | ✓ 874ms | http |
| 115.231.181.40:8128 | ✓ 1983ms | ✓ 874ms | ✓ 1037ms | ✓ 1877ms | ✓ 926ms | http |
| 107.175.85.198:1080 | ✓ 838ms | 否 | ✓ 1823ms | ✓ 1728ms | ✓ 1850ms | http |
| 113.160.132.26:8080 | ✓ 1639ms | 否 | ✓ 1520ms | ✓ 1384ms | ✓ 1520ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1102ms | ✓ 614ms | ✓ 826ms | ✓ 648ms | http |
| 212.58.132.5:8888 | ✓ 1519ms | 否 | ✓ 1565ms | ✓ 1423ms | ✓ 1190ms | http |
| 129.80.217.21:444 | ✓ 637ms | ✓ 1143ms | ✓ 1073ms | ✓ 1485ms | ✓ 1013ms | http |
| 168.110.52.228:3128 | ✓ 485ms | ✓ 1230ms | ✓ 440ms | ✓ 1827ms | ✓ 1649ms | http |
| 113.176.92.71:3128 | ✓ 1554ms | ✓ 1310ms | ✓ 1412ms | ✓ 1438ms | ✓ 908ms | http |
| 193.160.209.58:1080 | ✓ 1670ms | 否 | ✓ 1313ms | 否 | ✓ 1918ms | http |
| 91.242.229.129:8092 | ✓ 1592ms | ✓ 1696ms | 否 | 否 | ✓ 1368ms | http |
| 59.46.216.131:30001 | ✓ 1131ms | 否 | ✓ 1894ms | ✓ 1136ms | ✓ 1695ms | http |
| 103.82.23.118:5182 | ✓ 1259ms | ✓ 1951ms | ✓ 1310ms | ✓ 1815ms | ✓ 1586ms | http |
| 34.101.184.164:3128 | ✓ 1736ms | 否 | ✓ 1492ms | ✓ 1658ms | ✓ 937ms | http |
| 8.154.21.175:3128 | ✓ 723ms | ✓ 892ms | ✓ 660ms | ✓ 875ms | ✓ 660ms | http |
| 128.199.116.219:9090 | ✓ 792ms | 否 | ✓ 1157ms | ✓ 994ms | ✓ 809ms | http |
| 128.199.113.85:9090 | ✓ 728ms | 否 | ✓ 1075ms | ✓ 1018ms | ✓ 786ms | http |
| 146.190.80.158:9090 | ✓ 706ms | 否 | ✓ 1227ms | ✓ 1335ms | 否 | http |
| 128.199.121.61:9090 | ✓ 705ms | 否 | ✓ 1463ms | ✓ 1019ms | 否 | http |
| 152.42.170.187:9090 | ✓ 707ms | 否 | ✓ 1093ms | ✓ 1035ms | 否 | http |
| 128.199.254.13:9090 | ✓ 728ms | 否 | ✓ 1790ms | 否 | ✓ 886ms | http |
| 45.59.122.132:80 | ✓ 1761ms | ✓ 1912ms | ✓ 1265ms | ✓ 1546ms | ✓ 1875ms | http |
| 120.92.212.16:8890 | ✓ 1576ms | ✓ 1069ms | 否 | ✓ 1542ms | 否 | http |
| 148.230.4.241:999 | ✓ 529ms | ✓ 1530ms | ✓ 672ms | ✓ 1534ms | ✓ 1269ms | http |
| 8.219.97.248:80 | ✓ 1510ms | 否 | ✓ 1188ms | ✓ 1981ms | 否 | http |
| 120.92.212.16:7890 | ✓ 899ms | ✓ 1032ms | ✓ 1048ms | ✓ 1580ms | ✓ 868ms | http |
| 43.167.192.85:8080 | ✓ 1680ms | ✓ 1157ms | ✓ 501ms | ✓ 759ms | ✓ 801ms | http |
| 103.21.220.141:3128 | ✓ 673ms | 否 | ✓ 649ms | ✓ 823ms | ✓ 657ms | http |
| 47.101.159.19:8899 | ✓ 1786ms | ✓ 917ms | ✓ 806ms | ✓ 1029ms | ✓ 825ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 939ms | ✓ 1904ms | ✓ 933ms | http |
| 180.125.216.109:8118 | ✓ 849ms | 否 | 否 | ✓ 983ms | ✓ 683ms | http |
| 210.223.44.230:3128 | ✓ 1918ms | ✓ 1147ms | ✓ 1041ms | ✓ 1112ms | ✓ 673ms | http |
| 158.160.215.167:8125 | ✓ 1322ms | ✓ 1965ms | ✓ 1699ms | 否 | 否 | http |
| 158.160.215.167:8123 | ✓ 944ms | 否 | ✓ 1268ms | ✓ 1993ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1512ms | 否 | 否 | ✓ 1555ms | ✓ 1395ms | http |
| 34.71.229.255:3128 | ✓ 1232ms | 否 | ✓ 1641ms | ✓ 1679ms | 否 | http |
| 47.74.226.8:5001 | 否 | ✓ 1322ms | ✓ 902ms | ✓ 1236ms | 否 | http |
| 45.125.67.37:8443 | ✓ 919ms | 否 | ✓ 1267ms | 否 | ✓ 1940ms | http |
| 121.230.9.33:1080 | ✓ 1190ms | ✓ 1257ms | 否 | ✓ 1203ms | ✓ 1038ms | http |
| 104.248.151.93:9090 | ✓ 938ms | 否 | ✓ 692ms | ✓ 1020ms | ✓ 940ms | http |
| 121.130.177.28:8888 | ✓ 1480ms | ✓ 1142ms | ✓ 1777ms | ✓ 1477ms | 否 | http |
| 5.252.33.13:2025 | ✓ 1549ms | 否 | ✓ 1524ms | 否 | ✓ 1989ms | http |
| 3.101.133.120:80 | ✓ 1237ms | ✓ 1730ms | ✓ 1429ms | ✓ 1168ms | ✓ 913ms | http |
| 154.27.196.238:999 | ✓ 1171ms | ✓ 1507ms | ✓ 1915ms | ✓ 1452ms | ✓ 1888ms | http |
| 42.114.173.42:1111 | ✓ 1387ms | 否 | ✓ 1773ms | ✓ 1517ms | ✓ 1353ms | http |
| 103.235.34.38:3888 | ✓ 1749ms | 否 | 否 | ✓ 1529ms | ✓ 1442ms | http |
| 190.102.9.23:1111 | ✓ 1454ms | 否 | ✓ 1809ms | 否 | ✓ 1949ms | http |
| 91.233.223.147:3128 | ✓ 1870ms | 否 | ✓ 1645ms | 否 | ✓ 1971ms | http |
| 166.88.55.83:7890 | ✓ 666ms | ✓ 1148ms | ✓ 758ms | ✓ 783ms | ✓ 746ms | http |
| 223.16.170.103:80 | ✓ 861ms | 否 | ✓ 816ms | ✓ 1017ms | ✓ 1025ms | http |
| 45.129.141.143:3128 | ✓ 936ms | 否 | ✓ 1952ms | 否 | ✓ 1989ms | http |
| 52.59.51.29:10797 | ✓ 1539ms | 否 | ✓ 1877ms | 否 | ✓ 1536ms | http |
| 63.179.134.206:40083 | ✓ 1580ms | 否 | ✓ 1982ms | 否 | ✓ 1736ms | http |
| 3.71.26.7:8081 | ✓ 1554ms | 否 | ✓ 1866ms | ✓ 1900ms | 否 | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 937ms | ✓ 1083ms | ✓ 908ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 855ms | ✓ 918ms | ✓ 957ms | http |
| 84.47.150.125:1080 | ✓ 1929ms | 否 | ✓ 1504ms | 否 | ✓ 1548ms | http |
| 121.230.8.136:1080 | ✓ 972ms | ✓ 1303ms | ✓ 875ms | ✓ 1344ms | ✓ 1033ms | http |
| 121.230.8.133:1080 | ✓ 813ms | ✓ 1027ms | ✓ 729ms | ✓ 1374ms | 否 | http |
| 121.230.8.181:1080 | ✓ 915ms | ✓ 1605ms | ✓ 726ms | ✓ 1189ms | ✓ 1054ms | http |
| 101.32.243.189:80 | ✓ 1702ms | ✓ 1341ms | ✓ 1391ms | ✓ 1635ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1665ms | 否 | ✓ 1897ms | 否 | ✓ 1464ms | http |
| 109.234.38.35:3128 | ✓ 1513ms | ✓ 1747ms | ✓ 1814ms | 否 | ✓ 1708ms | http |
| 175.194.173.105:3128 | ✓ 1510ms | ✓ 937ms | ✓ 945ms | ✓ 936ms | ✓ 1185ms | http |
| 123.30.234.153:1315 | ✓ 1695ms | ✓ 1519ms | ✓ 1231ms | ✓ 1412ms | ✓ 918ms | http |
| 112.163.160.93:3128 | 否 | ✓ 1004ms | ✓ 820ms | ✓ 1794ms | 否 | http |
| 121.230.9.26:1080 | ✓ 963ms | 否 | ✓ 1035ms | ✓ 1163ms | ✓ 848ms | http |
| 61.52.131.172:8443 | ✓ 668ms | 否 | ✓ 1297ms | ✓ 1081ms | ✓ 761ms | http |
| 38.75.82.221:999 | ✓ 929ms | ✓ 1589ms | ✓ 1359ms | 否 | 否 | http |
| 103.157.117.116:8080 | ✓ 1706ms | 否 | ✓ 1866ms | ✓ 1733ms | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
