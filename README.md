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

最后更新：2026-03-11 20:43:25 UTC（2026-03-12 04:43:25 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | ✓ 1507ms | 否 | ✓ 1086ms | ✓ 927ms | http |
| 45.136.131.47:8443 | ✓ 607ms | ✓ 816ms | ✓ 1744ms | ✓ 1016ms | ✓ 914ms | http |
| 45.136.131.63:8443 | ✓ 805ms | ✓ 1008ms | ✓ 1361ms | ✓ 919ms | ✓ 1281ms | http |
| 45.136.130.175:8443 | ✓ 981ms | 否 | ✓ 988ms | ✓ 918ms | ✓ 966ms | http |
| 91.107.141.42:8081 | ✓ 490ms | 否 | ✓ 1268ms | ✓ 1853ms | 否 | http |
| 185.115.74.185:8080 | ✓ 695ms | ✓ 1954ms | ✓ 1322ms | 否 | 否 | http |
| 171.251.172.78:5109 | ✓ 1739ms | 否 | ✓ 1843ms | 否 | ✓ 1634ms | http |
| 190.9.109.198:999 | ✓ 1103ms | ✓ 1331ms | ✓ 1116ms | ✓ 1218ms | ✓ 993ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1238ms | ✓ 1497ms | ✓ 1207ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1717ms | ✓ 1180ms | ✓ 858ms | http |
| 101.43.255.96:80 | ✓ 1306ms | ✓ 1474ms | ✓ 1243ms | ✓ 1454ms | ✓ 1226ms | http |
| 81.70.169.194:80 | ✓ 1154ms | ✓ 1731ms | ✓ 1179ms | ✓ 1501ms | ✓ 1133ms | http |
| 5.252.33.13:2025 | ✓ 1311ms | 否 | ✓ 1147ms | ✓ 1966ms | ✓ 1780ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 935ms | ✓ 1197ms | ✓ 1913ms | http |
| 45.136.130.191:8443 | 否 | 否 | ✓ 1589ms | ✓ 1386ms | ✓ 1057ms | http |
| 45.136.130.188:8443 | 否 | 否 | ✓ 1847ms | ✓ 1431ms | ✓ 821ms | http |
| 165.227.5.10:8888 | ✓ 608ms | ✓ 1541ms | ✓ 515ms | ✓ 1175ms | ✓ 880ms | http |
| 103.35.188.243:3128 | ✓ 192ms | ✓ 1924ms | 否 | ✓ 1285ms | ✓ 1070ms | http |
| 202.155.12.161:443 | ✓ 1664ms | 否 | ✓ 1286ms | ✓ 1158ms | ✓ 1242ms | http |
| 62.234.206.73:3128 | ✓ 1210ms | ✓ 1482ms | ✓ 1114ms | ✓ 1421ms | ✓ 1179ms | http |
| 115.231.181.40:8128 | ✓ 1047ms | ✓ 1326ms | ✓ 1121ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 914ms | ✓ 1716ms | ✓ 1542ms | ✓ 1627ms | ✓ 1229ms | http |
| 205.209.118.30:3138 | ✓ 1721ms | ✓ 1822ms | ✓ 1738ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 1163ms | ✓ 1683ms | 否 | ✓ 1055ms | ✓ 1116ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1479ms | ✓ 1530ms | ✓ 1177ms | http |
| 107.173.52.58:7890 | ✓ 1437ms | ✓ 1331ms | ✓ 778ms | 否 | ✓ 995ms | http |
| 113.160.132.26:8080 | ✓ 1707ms | ✓ 1579ms | ✓ 1464ms | ✓ 1445ms | ✓ 1176ms | http |
| 180.127.149.252:1080 | 否 | 否 | ✓ 1040ms | ✓ 1493ms | ✓ 1066ms | http |
| 47.77.193.180:1080 | ✓ 855ms | ✓ 902ms | ✓ 567ms | ✓ 1003ms | ✓ 694ms | http |
| 1.231.81.166:3128 | ✓ 1846ms | ✓ 1258ms | ✓ 1441ms | ✓ 1169ms | ✓ 909ms | http |
| 121.230.9.241:1080 | ✓ 1469ms | ✓ 1576ms | ✓ 1462ms | 否 | ✓ 1322ms | http |
| 116.80.96.111:3172 | ✓ 1812ms | 否 | ✓ 1696ms | 否 | ✓ 1808ms | http |
| 107.172.125.217:3128 | ✓ 442ms | 否 | ✓ 1055ms | ✓ 958ms | ✓ 734ms | http |
| 111.48.191.1:7890 | ✓ 941ms | ✓ 1183ms | ✓ 992ms | ✓ 1165ms | ✓ 916ms | http |
| 168.235.110.63:3128 | ✓ 709ms | ✓ 1058ms | ✓ 511ms | 否 | 否 | http |
| 152.42.213.210:443 | ✓ 1569ms | 否 | ✓ 1323ms | ✓ 1451ms | ✓ 1000ms | http |
| 88.80.150.82:8080 | ✓ 955ms | 否 | 否 | ✓ 1941ms | ✓ 1578ms | https |
| 178.236.245.17:3128 | ✓ 1388ms | ✓ 1679ms | ✓ 1019ms | 否 | ✓ 1978ms | http |
| 180.103.19.53:1080 | 否 | 否 | ✓ 1672ms | ✓ 1598ms | ✓ 1268ms | http |
| 45.136.198.40:3128 | ✓ 1051ms | 否 | ✓ 1717ms | 否 | ✓ 1815ms | http |
| 152.42.213.210:8080 | ✓ 1832ms | 否 | ✓ 1180ms | ✓ 1515ms | ✓ 1023ms | http |
| 194.213.18.200:443 | ✓ 1496ms | 否 | ✓ 534ms | ✓ 1094ms | ✓ 1949ms | http |
| 95.3.9.78:3128 | ✓ 1841ms | 否 | 否 | ✓ 1531ms | ✓ 1684ms | http |
| 95.3.9.78:8080 | ✓ 1020ms | 否 | ✓ 1426ms | ✓ 1561ms | ✓ 1212ms | http |
| 106.117.208.101:7890 | ✓ 1335ms | 否 | ✓ 1783ms | 否 | ✓ 1594ms | http |
| 185.191.236.162:3128 | ✓ 1519ms | 否 | 否 | ✓ 1920ms | ✓ 1966ms | http |
| 45.140.147.82:1081 | ✓ 1034ms | 否 | ✓ 1060ms | ✓ 1669ms | ✓ 1262ms | http |
| 103.39.51.190:8080 | ✓ 1922ms | 否 | 否 | ✓ 1823ms | ✓ 1669ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1675ms | 否 | ✓ 1500ms | ✓ 1372ms | http |
| 45.136.130.223:8443 | 否 | ✓ 855ms | ✓ 946ms | ✓ 867ms | ✓ 738ms | http |
| 39.104.201.40:7890 | ✓ 918ms | 否 | ✓ 1390ms | ✓ 1241ms | ✓ 938ms | http |
| 46.183.25.8:443 | ✓ 1685ms | 否 | 否 | ✓ 1741ms | ✓ 1394ms | http |
| 117.18.20.102:8081 | 否 | 否 | ✓ 1672ms | ✓ 1781ms | ✓ 1780ms | http |
| 61.52.131.172:8443 | ✓ 943ms | ✓ 1282ms | ✓ 991ms | ✓ 1162ms | ✓ 916ms | http |
| 8.219.97.248:80 | ✓ 1978ms | 否 | ✓ 1342ms | ✓ 1990ms | 否 | http |
| 223.16.170.103:3128 | ✓ 1139ms | ✓ 1868ms | ✓ 1400ms | ✓ 1412ms | ✓ 1355ms | http |
| 116.80.96.100:3172 | ✓ 1716ms | 否 | ✓ 1682ms | 否 | ✓ 1808ms | http |
| 178.236.245.59:3128 | 否 | 否 | ✓ 1715ms | ✓ 1527ms | ✓ 1374ms | http |
| 34.101.184.164:3128 | ✓ 1757ms | 否 | ✓ 1548ms | 否 | ✓ 1914ms | http |

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
