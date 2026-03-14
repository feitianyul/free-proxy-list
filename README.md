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

最后更新：2026-03-14 20:31:22 UTC（2026-03-15 04:31:22 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.45.251.242:8888 | ✓ 947ms | ✓ 1787ms | 否 | ✓ 1832ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1567ms | ✓ 1588ms | ✓ 1607ms | ✓ 1466ms | 否 | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1134ms | ✓ 1382ms | ✓ 1445ms | http |
| 45.167.124.52:8080 | ✓ 1064ms | ✓ 1864ms | ✓ 1996ms | ✓ 1842ms | ✓ 1383ms | http |
| 91.233.223.147:3128 | ✓ 809ms | 否 | ✓ 761ms | ✓ 1925ms | ✓ 1481ms | http |
| 165.227.5.10:8888 | ✓ 899ms | 否 | ✓ 402ms | ✓ 1116ms | ✓ 849ms | http |
| 59.46.216.131:30001 | ✓ 1115ms | ✓ 1633ms | ✓ 1268ms | ✓ 1601ms | ✓ 1236ms | http |
| 45.136.131.42:8447 | ✓ 641ms | ✓ 1267ms | ✓ 1241ms | ✓ 1034ms | ✓ 712ms | http |
| 38.145.203.135:8443 | ✓ 639ms | ✓ 1448ms | ✓ 1063ms | ✓ 936ms | ✓ 986ms | http |
| 38.145.218.101:8447 | 否 | ✓ 1347ms | ✓ 802ms | ✓ 950ms | ✓ 906ms | http |
| 101.43.127.100:8877 | ✓ 999ms | ✓ 1277ms | ✓ 995ms | ✓ 1384ms | ✓ 1078ms | http |
| 81.70.169.194:80 | ✓ 1137ms | ✓ 1556ms | ✓ 1245ms | ✓ 1393ms | ✓ 1251ms | http |
| 101.43.255.96:80 | ✓ 1242ms | ✓ 1462ms | ✓ 1228ms | ✓ 1583ms | ✓ 1221ms | http |
| 210.77.29.245:7890 | ✓ 1248ms | ✓ 1847ms | 否 | ✓ 1612ms | ✓ 1047ms | http |
| 45.136.130.223:8443 | ✓ 407ms | 否 | ✓ 687ms | ✓ 1036ms | ✓ 709ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1903ms | ✓ 1183ms | 否 | ✓ 927ms | http |
| 85.198.96.242:3128 | ✓ 518ms | ✓ 1639ms | ✓ 1960ms | 否 | 否 | http |
| 38.145.218.82:8443 | ✓ 980ms | ✓ 910ms | ✓ 315ms | ✓ 1094ms | ✓ 817ms | http |
| 120.92.212.16:8890 | ✓ 1419ms | ✓ 1471ms | ✓ 1437ms | 否 | ✓ 1151ms | http |
| 185.115.74.185:8080 | ✓ 1344ms | ✓ 1735ms | ✓ 1911ms | 否 | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1211ms | ✓ 1305ms | ✓ 1375ms | http |
| 106.14.203.63:3333 | ✓ 1006ms | ✓ 1349ms | ✓ 1726ms | ✓ 1308ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1772ms | ✓ 1114ms | ✓ 1450ms | 否 | http |
| 37.187.109.70:10111 | ✓ 1655ms | ✓ 1962ms | ✓ 1014ms | ✓ 1589ms | ✓ 1188ms | http |
| 45.207.200.120:1080 | ✓ 1876ms | 否 | ✓ 1487ms | ✓ 1071ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1951ms | ✓ 1714ms | ✓ 1681ms | ✓ 1961ms | 否 | http |
| 35.225.22.61:80 | ✓ 1015ms | ✓ 1968ms | ✓ 832ms | ✓ 1066ms | ✓ 829ms | http |
| 205.209.118.30:3138 | ✓ 1039ms | ✓ 1659ms | ✓ 563ms | ✓ 1959ms | 否 | http |
| 43.167.227.161:1080 | ✓ 677ms | ✓ 1763ms | 否 | 否 | ✓ 1474ms | http |
| 45.136.130.214:8443 | ✓ 550ms | ✓ 864ms | ✓ 702ms | ✓ 1965ms | 否 | http |
| 45.136.130.217:8443 | ✓ 575ms | ✓ 1430ms | ✓ 1387ms | 否 | 否 | http |
| 95.3.9.78:8080 | ✓ 932ms | 否 | ✓ 1059ms | ✓ 1732ms | ✓ 1435ms | http |
| 168.235.110.63:3128 | ✓ 220ms | ✓ 940ms | ✓ 782ms | ✓ 916ms | ✓ 728ms | http |
| 88.80.150.82:8080 | ✓ 975ms | ✓ 1877ms | 否 | ✓ 1580ms | ✓ 1190ms | https |
| 91.107.148.58:53967 | ✓ 1006ms | ✓ 1819ms | 否 | 否 | ✓ 1927ms | http |
| 121.230.8.62:1080 | ✓ 1255ms | ✓ 1810ms | ✓ 1423ms | ✓ 1747ms | ✓ 1304ms | http |
| 45.136.130.216:8443 | ✓ 961ms | ✓ 936ms | ✓ 914ms | ✓ 1084ms | ✓ 799ms | http |
| 45.136.130.215:8443 | ✓ 963ms | ✓ 861ms | ✓ 989ms | ✓ 1115ms | ✓ 707ms | http |
| 38.145.218.235:8443 | ✓ 298ms | ✓ 925ms | ✓ 289ms | ✓ 883ms | ✓ 693ms | http |
| 34.101.184.164:3128 | ✓ 1150ms | 否 | ✓ 1831ms | ✓ 1862ms | ✓ 1186ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 871ms | ✓ 1337ms | ✓ 828ms | http |
| 164.90.155.209:3128 | ✓ 483ms | ✓ 1102ms | ✓ 1350ms | ✓ 935ms | ✓ 711ms | http |
| 128.199.120.45:9090 | ✓ 1643ms | 否 | ✓ 1388ms | ✓ 1582ms | ✓ 1451ms | http |
| 103.113.70.189:1081 | ✓ 1020ms | ✓ 893ms | ✓ 844ms | ✓ 1079ms | ✓ 806ms | http |
| 129.213.162.27:17777 | ✓ 1224ms | ✓ 1312ms | ✓ 1442ms | ✓ 1674ms | ✓ 1213ms | http |
| 62.113.119.14:8080 | ✓ 530ms | ✓ 1550ms | ✓ 561ms | ✓ 1433ms | 否 | http |
| 95.3.9.78:3128 | ✓ 1088ms | ✓ 1656ms | ✓ 1013ms | ✓ 1533ms | ✓ 1170ms | http |
| 1.234.153.14:80 | 否 | ✓ 1289ms | ✓ 994ms | ✓ 1357ms | ✓ 881ms | http |
| 2.56.122.146:10808 | 否 | 否 | ✓ 646ms | ✓ 1691ms | ✓ 853ms | http |
| 45.136.131.39:8443 | ✓ 1534ms | ✓ 1402ms | ✓ 1250ms | ✓ 1510ms | ✓ 1133ms | http |
| 45.136.130.162:8443 | ✓ 355ms | ✓ 1457ms | ✓ 773ms | ✓ 938ms | ✓ 723ms | http |
| 45.136.130.155:8443 | ✓ 355ms | ✓ 1425ms | ✓ 815ms | ✓ 1014ms | ✓ 740ms | http |
| 38.145.208.239:8443 | ✓ 423ms | ✓ 844ms | ✓ 1189ms | ✓ 1692ms | ✓ 957ms | http |
| 45.136.130.156:8443 | ✓ 322ms | ✓ 1552ms | ✓ 317ms | ✓ 962ms | ✓ 767ms | http |
| 45.136.131.32:8443 | ✓ 323ms | ✓ 911ms | ✓ 744ms | ✓ 1177ms | ✓ 779ms | http |
| 45.136.131.30:8447 | ✓ 353ms | ✓ 935ms | ✓ 690ms | ✓ 1136ms | ✓ 832ms | http |
| 45.136.131.33:8443 | ✓ 410ms | ✓ 1390ms | ✓ 337ms | ✓ 1042ms | ✓ 790ms | http |
| 45.136.131.31:8443 | ✓ 301ms | ✓ 883ms | ✓ 803ms | ✓ 1174ms | ✓ 814ms | http |
| 45.136.131.35:8443 | ✓ 293ms | ✓ 879ms | ✓ 805ms | ✓ 1171ms | ✓ 782ms | http |
| 150.107.140.238:3128 | ✓ 1942ms | 否 | 否 | ✓ 1686ms | ✓ 1453ms | http |
| 150.249.255.91:3128 | ✓ 1161ms | ✓ 1831ms | ✓ 685ms | ✓ 1099ms | ✓ 842ms | http |
| 38.145.203.162:8443 | 否 | ✓ 1986ms | ✓ 287ms | ✓ 1135ms | 否 | http |
| 38.145.208.93:8443 | 否 | ✓ 1975ms | ✓ 285ms | ✓ 1454ms | 否 | http |
| 38.145.208.94:8443 | 否 | ✓ 1986ms | ✓ 322ms | ✓ 1974ms | 否 | http |
| 45.136.130.161:8443 | ✓ 573ms | ✓ 946ms | ✓ 775ms | ✓ 976ms | ✓ 723ms | http |
| 159.223.42.219:3128 | ✓ 1405ms | 否 | ✓ 1536ms | ✓ 1277ms | ✓ 1012ms | http |
| 45.136.131.28:8447 | ✓ 1050ms | ✓ 1227ms | ✓ 612ms | ✓ 1283ms | ✓ 1610ms | http |
| 104.243.46.122:3128 | ✓ 1707ms | ✓ 1245ms | ✓ 1537ms | ✓ 1189ms | ✓ 998ms | http |
| 61.52.131.172:8443 | ✓ 1116ms | ✓ 1324ms | ✓ 1092ms | ✓ 1365ms | ✓ 1091ms | http |
| 8.140.104.98:3128 | ✓ 1084ms | ✓ 1394ms | ✓ 1150ms | ✓ 1552ms | ✓ 1145ms | http |
| 14.225.212.37:7890 | ✓ 1014ms | 否 | ✓ 1051ms | 否 | ✓ 1095ms | http |
| 106.117.208.101:7890 | ✓ 1191ms | ✓ 1610ms | ✓ 1722ms | ✓ 1523ms | ✓ 1301ms | http |
| 138.124.53.221:443 | ✓ 429ms | ✓ 1569ms | ✓ 1465ms | ✓ 1818ms | 否 | http |

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
