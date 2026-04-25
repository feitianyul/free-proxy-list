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

最后更新：2026-04-25 08:13:58 UTC（2026-04-25 16:13:58 UTC+8）

**代理总数：48**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1998ms | 否 | ✓ 1945ms | ✓ 1792ms | ✓ 1463ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1626ms | ✓ 895ms | ✓ 1105ms | ✓ 1599ms | http |
| 2.27.54.161:1080 | ✓ 1498ms | 否 | ✓ 1428ms | ✓ 1993ms | ✓ 1500ms | http |
| 218.108.131.186:17890 | ✓ 1879ms | ✓ 1284ms | ✓ 1242ms | ✓ 1589ms | 否 | http |
| 43.165.195.107:3128 | ✓ 1709ms | ✓ 1766ms | ✓ 1383ms | ✓ 1387ms | ✓ 1109ms | http |
| 59.46.216.131:30001 | ✓ 1172ms | ✓ 1567ms | ✓ 1222ms | ✓ 1528ms | 否 | http |
| 45.153.231.229:8080 | ✓ 774ms | 否 | ✓ 1525ms | ✓ 1898ms | ✓ 1373ms | http |
| 38.180.192.119:3128 | ✓ 995ms | ✓ 936ms | 否 | ✓ 1069ms | ✓ 708ms | http |
| 161.35.181.96:999 | 否 | ✓ 1715ms | ✓ 500ms | ✓ 1102ms | ✓ 992ms | http |
| 177.93.132.244:3128 | ✓ 1315ms | 否 | ✓ 1783ms | 否 | ✓ 1696ms | http |
| 152.70.91.193:40000 | ✓ 1558ms | ✓ 1931ms | ✓ 1424ms | ✓ 1558ms | 否 | http |
| 166.88.61.54:8000 | ✓ 1218ms | ✓ 1414ms | ✓ 1134ms | ✓ 1049ms | ✓ 1179ms | http |
| 91.233.223.147:3128 | ✓ 1462ms | 否 | ✓ 1435ms | 否 | ✓ 1847ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1495ms | ✓ 1518ms | ✓ 1320ms | http |
| 94.131.106.231:1081 | 否 | ✓ 1725ms | ✓ 1724ms | ✓ 1594ms | 否 | http |
| 206.206.126.177:2412 | ✓ 1530ms | 否 | ✓ 1148ms | ✓ 1202ms | ✓ 952ms | http |
| 120.92.212.16:8890 | ✓ 1111ms | 否 | ✓ 1360ms | ✓ 1373ms | ✓ 1693ms | http |
| 84.47.150.125:1080 | ✓ 1984ms | 否 | ✓ 1903ms | ✓ 1902ms | 否 | http |
| 1.231.81.166:3128 | ✓ 865ms | 否 | ✓ 1027ms | ✓ 1396ms | ✓ 1191ms | http |
| 38.180.2.107:3128 | ✓ 889ms | ✓ 1601ms | 否 | 否 | ✓ 1863ms | http |
| 159.223.225.118:8888 | ✓ 722ms | 否 | ✓ 1543ms | ✓ 1610ms | ✓ 1267ms | http |
| 86.104.74.110:1081 | 否 | ✓ 1963ms | ✓ 1296ms | ✓ 1703ms | ✓ 1350ms | http |
| 114.231.72.27:1080 | ✓ 1242ms | 否 | ✓ 1555ms | 否 | ✓ 1487ms | http |
| 45.10.70.209:8888 | ✓ 1054ms | ✓ 1455ms | ✓ 489ms | ✓ 1093ms | ✓ 886ms | http |
| 45.186.6.104:3128 | ✓ 1713ms | ✓ 1895ms | ✓ 1932ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 442ms | 否 | ✓ 1283ms | 否 | ✓ 909ms | http |
| 37.187.109.70:10111 | ✓ 1245ms | ✓ 1679ms | ✓ 1723ms | 否 | 否 | http |
| 47.84.59.16:1080 | ✓ 1190ms | ✓ 1851ms | ✓ 1055ms | ✓ 1282ms | ✓ 1001ms | http |
| 34.71.229.255:3128 | ✓ 1660ms | 否 | ✓ 1453ms | ✓ 1457ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1458ms | 否 | ✓ 990ms | 否 | ✓ 1369ms | http |
| 130.61.174.200:1080 | ✓ 1084ms | ✓ 1567ms | 否 | ✓ 1755ms | 否 | http |
| 183.232.248.73:7890 | ✓ 1195ms | 否 | ✓ 1496ms | ✓ 1264ms | ✓ 1951ms | http |
| 120.92.212.16:7890 | ✓ 1194ms | 否 | 否 | ✓ 1485ms | ✓ 1586ms | http |
| 91.217.81.131:1080 | ✓ 1297ms | ✓ 1944ms | ✓ 1659ms | 否 | ✓ 1396ms | http |
| 45.129.141.143:3128 | ✓ 677ms | ✓ 1690ms | ✓ 1729ms | ✓ 1905ms | ✓ 1571ms | http |
| 47.84.76.30:1080 | 否 | ✓ 1927ms | ✓ 935ms | ✓ 1281ms | ✓ 1037ms | http |
| 47.84.73.61:1080 | ✓ 893ms | ✓ 1908ms | ✓ 891ms | ✓ 1282ms | ✓ 1027ms | http |
| 91.99.15.45:2095 | ✓ 599ms | ✓ 1553ms | ✓ 998ms | ✓ 1969ms | 否 | http |
| 64.188.77.26:3128 | ✓ 1977ms | ✓ 1788ms | ✓ 1876ms | 否 | 否 | http |
| 54.37.72.89:80 | ✓ 1589ms | ✓ 1711ms | ✓ 1947ms | 否 | 否 | http |
| 121.230.9.160:1080 | 否 | ✓ 1569ms | ✓ 1259ms | 否 | ✓ 1348ms | http |
| 120.92.108.86:7890 | ✓ 1941ms | 否 | 否 | ✓ 1991ms | ✓ 1527ms | http |
| 212.58.132.5:8888 | ✓ 1373ms | 否 | ✓ 1842ms | ✓ 1776ms | 否 | http |
| 43.133.44.89:8888 | ✓ 1973ms | ✓ 1641ms | ✓ 1056ms | ✓ 1209ms | ✓ 976ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1507ms | 否 | ✓ 1604ms | ✓ 1366ms | http |
| 121.230.8.55:1080 | ✓ 1265ms | 否 | ✓ 1348ms | ✓ 1709ms | ✓ 1368ms | http |
| 94.131.122.128:1081 | 否 | 否 | ✓ 1381ms | ✓ 1452ms | ✓ 1399ms | http |
| 171.234.134.26:6616 | ✓ 1942ms | 否 | ✓ 1674ms | ✓ 1742ms | ✓ 1663ms | http |

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
