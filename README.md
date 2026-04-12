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

最后更新：2026-04-12 16:33:45 UTC（2026-04-13 00:33:45 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 798ms | ✓ 968ms | ✓ 826ms | ✓ 1549ms | ✓ 955ms | http |
| 147.161.210.140:8800 | ✓ 658ms | ✓ 1552ms | ✓ 639ms | ✓ 1280ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1715ms | 否 | ✓ 1160ms | ✓ 1050ms | ✓ 1052ms | http |
| 113.160.132.26:8080 | ✓ 1458ms | 否 | ✓ 1276ms | ✓ 1181ms | ✓ 1141ms | http |
| 36.103.198.235:7890 | ✓ 1875ms | ✓ 1883ms | ✓ 1398ms | ✓ 1373ms | 否 | http |
| 35.225.22.61:80 | ✓ 547ms | 否 | 否 | ✓ 1323ms | ✓ 1085ms | http |
| 36.141.21.200:7890 | ✓ 911ms | ✓ 1092ms | ✓ 958ms | ✓ 1221ms | ✓ 971ms | http |
| 223.84.151.86:30005 | ✓ 1180ms | ✓ 1179ms | ✓ 951ms | ✓ 1275ms | ✓ 1141ms | http |
| 1.231.81.166:3128 | ✓ 848ms | 否 | ✓ 1282ms | ✓ 1141ms | ✓ 976ms | http |
| 167.103.144.127:8800 | ✓ 1355ms | 否 | ✓ 1311ms | ✓ 1472ms | ✓ 1825ms | http |
| 167.103.34.108:8800 | ✓ 1479ms | 否 | ✓ 1322ms | ✓ 1369ms | ✓ 1216ms | http |
| 103.229.127.31:7890 | ✓ 1547ms | 否 | ✓ 1436ms | ✓ 1570ms | ✓ 1043ms | http |
| 171.227.167.109:1006 | 否 | 否 | ✓ 1407ms | ✓ 1361ms | ✓ 1413ms | http |
| 20.210.39.153:8561 | ✓ 1140ms | ✓ 1584ms | ✓ 1269ms | ✓ 1703ms | ✓ 1507ms | http |
| 20.78.26.206:8561 | ✓ 1167ms | ✓ 1693ms | ✓ 1200ms | ✓ 1666ms | ✓ 1509ms | http |
| 20.78.118.91:8561 | ✓ 1164ms | 否 | ✓ 1144ms | ✓ 1646ms | ✓ 1420ms | http |
| 45.167.125.21:999 | ✓ 989ms | 否 | ✓ 1400ms | 否 | ✓ 1702ms | http |
| 91.233.223.147:3128 | ✓ 1612ms | 否 | ✓ 1198ms | 否 | ✓ 1839ms | http |
| 167.103.31.122:8800 | ✓ 1307ms | 否 | ✓ 1295ms | 否 | ✓ 1504ms | http |
| 120.92.108.86:7890 | ✓ 1182ms | 否 | ✓ 1892ms | 否 | ✓ 1393ms | http |
| 45.167.124.52:8080 | ✓ 1100ms | 否 | ✓ 724ms | ✓ 1761ms | ✓ 1395ms | http |
| 212.58.132.5:8888 | ✓ 1239ms | 否 | ✓ 1399ms | ✓ 1479ms | ✓ 1229ms | http |
| 139.159.99.242:8080 | ✓ 766ms | ✓ 974ms | ✓ 781ms | ✓ 1015ms | ✓ 833ms | http |
| 8.219.195.129:1080 | ✓ 679ms | 否 | ✓ 930ms | ✓ 993ms | ✓ 794ms | http |
| 147.161.239.240:8800 | ✓ 914ms | 否 | ✓ 1603ms | 否 | ✓ 1569ms | http |
| 5.104.87.17:8051 | ✓ 933ms | 否 | ✓ 1136ms | 否 | ✓ 1531ms | http |
| 103.157.200.126:3128 | ✓ 1941ms | 否 | ✓ 1760ms | ✓ 1915ms | ✓ 1533ms | http |
| 140.238.254.5:8080 | ✓ 1522ms | 否 | ✓ 915ms | 否 | ✓ 1654ms | http |
| 121.230.8.220:1080 | ✓ 1491ms | ✓ 1406ms | ✓ 1415ms | ✓ 1717ms | 否 | http |
| 162.240.154.26:3128 | 否 | ✓ 1359ms | 否 | ✓ 1617ms | ✓ 1108ms | http |
| 5.104.87.17:8050 | ✓ 611ms | ✓ 1642ms | 否 | ✓ 1279ms | ✓ 716ms | http |
| 79.132.136.58:3128 | ✓ 762ms | ✓ 1879ms | ✓ 1242ms | ✓ 1564ms | ✓ 1207ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1073ms | ✓ 803ms | 否 | ✓ 950ms | http |
| 147.45.136.99:3128 | ✓ 994ms | 否 | ✓ 817ms | ✓ 1645ms | ✓ 1573ms | http |
| 210.223.44.230:3128 | ✓ 1130ms | ✓ 1907ms | ✓ 1919ms | ✓ 1412ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1685ms | ✓ 1093ms | ✓ 896ms | ✓ 1193ms | ✓ 975ms | http |
| 5.196.101.18:3128 | ✓ 1458ms | ✓ 1823ms | ✓ 1037ms | 否 | 否 | http |
| 51.79.207.21:8080 | ✓ 1100ms | 否 | ✓ 1556ms | ✓ 1907ms | ✓ 1227ms | http |
| 20.2.83.243:3128 | ✓ 788ms | ✓ 1786ms | ✓ 915ms | ✓ 1591ms | 否 | http |
| 119.92.70.232:8082 | ✓ 1334ms | 否 | ✓ 1647ms | 否 | ✓ 1585ms | http |
| 104.248.149.186:3128 | 否 | 否 | ✓ 686ms | ✓ 1026ms | ✓ 801ms | http |
| 62.113.119.14:8080 | ✓ 1734ms | ✓ 1814ms | ✓ 1110ms | 否 | 否 | http |
| 43.165.195.107:3128 | ✓ 1775ms | 否 | 否 | ✓ 1133ms | ✓ 909ms | http |
| 222.228.171.92:8080 | ✓ 1513ms | 否 | ✓ 1603ms | 否 | ✓ 1594ms | http |
| 103.252.93.142:3128 | ✓ 939ms | 否 | ✓ 1002ms | ✓ 1158ms | ✓ 936ms | http |
| 59.46.216.131:30001 | ✓ 1008ms | 否 | ✓ 1501ms | ✓ 1983ms | ✓ 1175ms | http |
| 8.219.97.248:80 | ✓ 1556ms | 否 | ✓ 1071ms | ✓ 1383ms | 否 | http |
| 107.172.102.234:40621 | ✓ 290ms | ✓ 1514ms | ✓ 1755ms | ✓ 1027ms | ✓ 593ms | http |
| 150.249.255.91:3128 | ✓ 624ms | 否 | ✓ 626ms | ✓ 807ms | ✓ 800ms | http |
| 24.144.86.173:1080 | 否 | 否 | ✓ 984ms | ✓ 930ms | ✓ 620ms | http |
| 34.101.184.164:3128 | ✓ 1726ms | 否 | ✓ 1251ms | 否 | ✓ 1600ms | http |
| 45.140.147.155:1081 | ✓ 1293ms | ✓ 1442ms | ✓ 627ms | ✓ 1433ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1767ms | ✓ 1975ms | 否 | 否 | ✓ 1324ms | http |
| 2.27.32.81:3128 | ✓ 1310ms | 否 | ✓ 1133ms | 否 | ✓ 1955ms | http |
| 130.61.30.221:8080 | ✓ 869ms | 否 | ✓ 1415ms | ✓ 1899ms | 否 | http |
| 46.30.46.133:3128 | ✓ 871ms | ✓ 1601ms | ✓ 747ms | ✓ 1826ms | 否 | http |
| 47.84.131.156:8100 | ✓ 882ms | ✓ 1646ms | ✓ 827ms | ✓ 1021ms | ✓ 826ms | http |
| 5.255.123.43:1080 | ✓ 678ms | ✓ 1319ms | ✓ 1147ms | 否 | 否 | http |
| 168.110.52.228:3128 | ✓ 1397ms | 否 | 否 | ✓ 828ms | ✓ 1791ms | http |
| 195.26.224.49:3128 | ✓ 1249ms | 否 | ✓ 1461ms | ✓ 1903ms | ✓ 1676ms | http |
| 61.52.131.172:8443 | ✓ 887ms | ✓ 1161ms | ✓ 881ms | ✓ 1123ms | ✓ 956ms | http |
| 103.113.70.189:1081 | ✓ 1534ms | 否 | ✓ 526ms | 否 | ✓ 902ms | http |

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
