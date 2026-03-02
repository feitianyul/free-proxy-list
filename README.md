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

最后更新：2026-03-02 17:54:59 UTC（2026-03-03 01:54:59 UTC+8）

**代理总数：51**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 266ms | 否 | ✓ 807ms | ✓ 1349ms | ✓ 851ms | http |
| 35.225.22.61:80 | ✓ 527ms | 否 | ✓ 433ms | 否 | ✓ 775ms | http |
| 103.84.95.54:7890 | ✓ 1016ms | 否 | ✓ 1294ms | ✓ 999ms | 否 | http |
| 217.76.245.80:999 | ✓ 841ms | ✓ 1535ms | ✓ 1295ms | 否 | ✓ 1171ms | http |
| 45.22.209.157:8888 | ✓ 886ms | 否 | ✓ 1278ms | ✓ 1774ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1275ms | 否 | ✓ 1644ms | 否 | ✓ 1973ms | http |
| 165.227.5.10:8888 | ✓ 297ms | 否 | 否 | ✓ 1758ms | ✓ 881ms | http |
| 142.171.85.32:1080 | ✓ 410ms | ✓ 1461ms | ✓ 1074ms | ✓ 1032ms | ✓ 1136ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1346ms | ✓ 1825ms | ✓ 1205ms | http |
| 91.238.104.171:2023 | ✓ 1214ms | 否 | ✓ 1720ms | 否 | ✓ 1833ms | http |
| 81.70.169.194:80 | ✓ 1270ms | ✓ 1881ms | ✓ 1201ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1376ms | 否 | ✓ 1702ms | 否 | ✓ 1284ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1193ms | ✓ 1344ms | ✓ 1112ms | http |
| 195.123.209.48:3128 | ✓ 1196ms | 否 | ✓ 882ms | ✓ 1470ms | ✓ 1151ms | http |
| 45.125.67.37:8443 | ✓ 1505ms | 否 | 否 | ✓ 1481ms | ✓ 1534ms | http |
| 210.223.44.230:3128 | ✓ 1483ms | ✓ 1595ms | ✓ 1386ms | 否 | ✓ 1064ms | http |
| 5.75.196.26:40000 | ✓ 491ms | 否 | ✓ 1822ms | 否 | ✓ 1755ms | http |
| 138.124.53.25:7443 | ✓ 1736ms | ✓ 1945ms | ✓ 1772ms | ✓ 1922ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1181ms | 否 | ✓ 1851ms | 否 | ✓ 1712ms | http |
| 120.92.212.16:7890 | ✓ 1882ms | ✓ 1419ms | ✓ 1402ms | 否 | ✓ 1847ms | http |
| 45.88.0.116:3128 | ✓ 810ms | ✓ 1546ms | ✓ 1506ms | ✓ 1435ms | ✓ 963ms | http |
| 45.88.0.111:3128 | ✓ 827ms | 否 | ✓ 1053ms | ✓ 1427ms | ✓ 958ms | http |
| 45.88.0.114:3128 | ✓ 827ms | 否 | ✓ 1052ms | ✓ 1435ms | ✓ 960ms | http |
| 45.88.0.113:3128 | ✓ 809ms | 否 | ✓ 1053ms | ✓ 1435ms | ✓ 987ms | http |
| 77.83.203.6:443 | ✓ 1047ms | ✓ 1981ms | ✓ 1681ms | 否 | 否 | http |
| 77.83.203.5:443 | ✓ 1039ms | ✓ 1740ms | ✓ 1921ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 847ms | ✓ 1958ms | 否 | 否 | ✓ 1788ms | http |
| 45.88.0.115:3128 | ✓ 1139ms | 否 | 否 | ✓ 1846ms | ✓ 1344ms | http |
| 45.88.0.98:3128 | ✓ 672ms | 否 | ✓ 755ms | ✓ 1252ms | ✓ 974ms | http |
| 45.88.0.117:3128 | ✓ 669ms | 否 | ✓ 757ms | ✓ 1252ms | ✓ 985ms | http |
| 45.88.0.99:3128 | ✓ 673ms | 否 | ✓ 753ms | ✓ 1265ms | ✓ 979ms | http |
| 46.249.103.192:443 | ✓ 935ms | 否 | ✓ 1574ms | ✓ 1439ms | 否 | http |
| 45.136.198.40:3128 | ✓ 883ms | ✓ 1794ms | ✓ 1367ms | ✓ 1733ms | ✓ 1612ms | http |
| 95.85.252.153:21064 | ✓ 548ms | ✓ 1810ms | ✓ 1368ms | ✓ 1912ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1256ms | ✓ 1284ms | ✓ 1964ms | 否 | http |
| 107.173.111.110:7890 | ✓ 1020ms | 否 | ✓ 1357ms | ✓ 1760ms | ✓ 1317ms | http |
| 101.32.244.83:8080 | ✓ 1611ms | 否 | ✓ 1071ms | ✓ 1404ms | ✓ 1352ms | http |
| 121.43.196.210:8222 | ✓ 1025ms | ✓ 1232ms | ✓ 1006ms | ✓ 1229ms | ✓ 977ms | http |
| 121.43.196.213:8222 | ✓ 1054ms | ✓ 1219ms | ✓ 1034ms | ✓ 1240ms | ✓ 1008ms | http |
| 114.55.226.123:10086 | ✓ 1094ms | ✓ 1870ms | ✓ 1118ms | ✓ 1378ms | ✓ 1169ms | http |
| 62.113.119.14:8080 | ✓ 1108ms | 否 | ✓ 716ms | ✓ 1483ms | ✓ 1115ms | http |
| 115.76.5.32:10010 | ✓ 1854ms | 否 | ✓ 1617ms | ✓ 1824ms | ✓ 1919ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1092ms | ✓ 1322ms | ✓ 1047ms | http |
| 45.186.6.104:3128 | ✓ 1237ms | ✓ 1964ms | ✓ 1708ms | 否 | 否 | http |
| 223.16.170.103:3128 | ✓ 1037ms | 否 | ✓ 1153ms | 否 | ✓ 1248ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1387ms | ✓ 1544ms | ✓ 1501ms | http |
| 158.160.215.167:8127 | ✓ 1626ms | 否 | ✓ 1637ms | 否 | ✓ 1315ms | http |
| 115.76.5.32:10007 | ✓ 1862ms | 否 | ✓ 1796ms | ✓ 1738ms | ✓ 1459ms | http |
| 45.177.178.242:999 | ✓ 1865ms | ✓ 1378ms | ✓ 1014ms | 否 | 否 | http |
| 115.76.5.32:10009 | ✓ 1860ms | 否 | 否 | ✓ 1740ms | ✓ 1478ms | http |
| 103.179.218.14:8000 | ✓ 1852ms | 否 | ✓ 1362ms | ✓ 1516ms | ✓ 1492ms | http |

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
